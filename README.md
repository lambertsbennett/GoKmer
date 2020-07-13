# GoKmer

In the work we are currently undertaking, trying to apply NLP techniques to sequence data, I found that I often needed to expand sequences into k-mers.
This is not particularly an enriching experience, so I worked on a small tool that concurrently expands sequences into k-mers of a given size. This tool operates
quite rapidly and can process large assemblies in seconds on an 8 core laptop.

## Usage
```
gokmer -file PATH_TO_SEQS -n NUM_PROCESSORS -k-mer-size K_MER_SIZE -o OUTPUT_FILE
```

**INPUTS:**
- Sequence Fasta file.

**OUTPUTS:**
- Parquet file with sequence ID and k-mers.
