// https://github.com/etcd-io/bbolt#reading-the-source
// https://www.google.com/search?q=golang+closure

package main

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
