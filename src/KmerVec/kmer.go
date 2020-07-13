package KmerVec

import (
	"fmt"
	"github.com/xitongsys/parquet-go/parquet"
	"log"
	"strings"
	"sync"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

func GetKmers(k int,wg *sync.WaitGroup, in chan Sequence,lsp *SequenceCollection){
	for s := range in {
		var b strings.Builder
		for j := 0; j < len(s.Seq)-(k-1); j++ {
			fmt.Fprint(&b, s.Seq[j : j+k], " ")
		}

	rs := NewSequence()
	rs.Header = s.Header
	rs.Kmers = b.String()
	lsp.Append(*rs)
	}
	wg.Done()
}

func (sc *SequenceCollection) ToParquet(fname string){
	type tmpseq struct {
		Header    string  `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
		Kmers     string   `parquet:"name=kmers, type=UTF8, encoding=PLAIN_DICTIONARY"`
	}

	fw, err := local.NewLocalFileWriter(fname)
	if err != nil {
		log.Println("Can't open file", err)
		return
	}
	pw, err := writer.NewParquetWriter(fw, new(tmpseq),4)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}
	pw.RowGroupSize = 5 * 1024 * 1024 //5M
	pw.CompressionType = parquet.CompressionCodec_GZIP

	for _,s := range sc.Items {
		seq := tmpseq{
			Header: s.Header,
			Kmers:  s.Kmers,
		}

		if err = pw.Write(seq); err != nil {
			log.Println("Write error", err)
		}
	}

	if err = pw.WriteStop(); err != nil {
		log.Println("WriteStop error", err)
		return
	}
log.Println("Write Finished")
fw.Close()
}