package cli

import (
	"bufio"
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	mrand "math/rand"

	golog "github.com/ipfs/go-log"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	net "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
)

func NewHost(port int, insecure bool, randseed uint64) (host.Host, error) {
	var r io.Reader

	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)

	if err != nil {
		return nil, err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
	}

	if insecure {
		opts = append(opts, libp2p.NoSecurity)
	}

	basicHost, err := libp2p.New(context.Background(), opts...)

	if err != nil {
		return nil, err
	}

	hostAddr, err := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

	addr := basicHost.Addrs()[0]

	fullAddr := addr.Encapsulate(hostAddr)

	if insecure {
		log.Printf("Now run \"./echo -l %d -d %s -insecure\" on a different terminal\n", listenPort+1, fullAddr)
	} else {
		log.Printf("Now run \"./echo -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
	}

	return basicHost, nil

}

func doEcho(s net.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	log.Printf("read: %s\n", str)
	_, err = s.Write([]byte(str))
	return err
}

func main() {
	golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info

	listenF := flag.Int("1", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	insecure := flag.Bool("insecure", false, "use an unencrypted connection")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	ha, err := NewHost(*listenF, *insecure, *seed)

	if err != nil {
		log.Fatal(err)
	}

	ha.SetStreamHandler("/echo/1.0.0", func(s net.Stream) {
		log.Println("Got a new stream")
		if err := doEcho(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	if *target == "" {
		log.Println("listening for connections")
		select {}
	}

	ipfsaddr, err := ma.NewMultiaddr(*target)

	if err != nil {
		log.Fatalln(err)
	}

	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)

	if err != nil {
		log.Fatalln(err)
	}

	peerid, err := peer.IDB58Decode(pid)

	if err != nil {
		log.Fatalln(err)
	}

	targetPeerAddr, _ := ma.NewMultiaddr(
		fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid))
	)

	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

	log.Println("opening stream")

	s, err := ha.NewStream(context.Background(), peerid, "/echo/1.0.0")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = s.Write([]byte("Hello, world!\n"))
	if err != nil {
		log.Fatalln(err)
	}

	out, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("read reply: %q\n", out)
}