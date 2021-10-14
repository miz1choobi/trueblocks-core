## chifra chunks

This tool is not yet ready for production use. Please return to this page later.

```[plaintext]
Purpose:
  Manage and investigate chunks and bloom filters.

Usage:
  chifra chunks [flags]

Flags:
  -l, --list      list the bloom and index hashes from local cache or IPFS
  -c, --check     check the validity of the chunk or bloom
  -e, --extract   show the contents of the chunk or bloom filters
  -s, --stats     for the --list option only, display statistics about each chunk or bloom

Global Flags:
  -x, --fmt string   export format, one of [none|json*|txt|csv|api]
  -h, --help         display this help screen
  -v, --verbose      enable verbose (increase detail with --log_level)
```

**Source code**: [`apps/chunkMan`](https://github.com/TrueBlocks/trueblocks-core/tree/master/src/apps/chunkMan)
