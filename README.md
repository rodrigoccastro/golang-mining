  1 - you need install and run app bitcoin core, because you need connected to blockchain
  2 - add this info in your file config of bitcoin core
  # Bitcoin Core configuration file
  
  # Run Bitcoin Core in server mode (required for RPC)
  server=1
  
  # Accept RPC connections from local clients only
  rpcbind=127.0.0.1
  
  # Port for RPC connections (default is 8332)
  rpcport=8332
  
  # Username and password for RPC authentication
  rpcuser=your_rpc_username
  rpcpassword=your_rpc_password
  
  # Specify the location of the data directory (optional)
  # datadir=/path/to/your/data/directory
  
  # Specify optional parameters, such as pruning, index, etc.
  prune=550
  # txindex=1
  
  # Add nodes to connect to (optional)
  # addnode=example.com
  
  walletbroadcast=0
  daemon=1
  rest=1
  zmqpubhashtx=tcp://127.0.0.1:28332
  rpcworkqueue=100
  maxwitnesssiz=20000
  
  3 - put your bitcoint address account in serviceGetDecodeAddress.go
