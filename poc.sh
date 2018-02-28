rm data.bolt
taster@power:~/go/src/gobtc|master⚡
⇒  go build
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc
Usage:
   printchain - Print all the blocks of the blockchain
   createblockchain -address ADDRESS -  Create a blockchain and send genesis block reward to ADDRESS
   getbalance -address ADDRESS - Get the balance of ADDRESS
   send -from FROM -to TO -amount AMOUNT - send AOUNT coins from FROM to TO
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc printchain
No existing blockchain found. Create one first.
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc create blockchain -address Jack
Usage:
   printchain - Print all the blocks of the blockchain
   createblockchain -address ADDRESS -  Create a blockchain and send genesis block reward to ADDRESS
   getbalance -address ADDRESS - Get the balance of ADDRESS
   send -from FROM -to TO -amount AMOUNT - send AOUNT coins from FROM to TO
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc createblockchain -address Jack
Mining block containing " [Transaction 9355a572b4e6218950a742097287942e104024585552e4356b2d144c5ce393bf:
  Input 0:
    TXID:
    Out:     -1
    Script:  Decentralization make more people have more right
  Output 0:
    Value:  50
    Script: Jack
] "


Done!
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc send -from Jack -to Ann -amount 10
Mining block containing " [Transaction 1256e1ca749b50a8d8bda37046b0ffceaf5f3bb2db019cff3141970fbd8bf405:
  Input 0:
    TXID:    9355a572b4e6218950a742097287942e104024585552e4356b2d144c5ce393bf
    Out:     0
    Script:  Jack
  Output 0:
    Value:  10
    Script: Ann
  Output 1:
    Value:  40
    Script: Jack
] "


Success
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc printchain
Timestamp: 1519799993
[Transaction 1256e1ca749b50a8d8bda37046b0ffceaf5f3bb2db019cff3141970fbd8bf405:
  Input 0:
    TXID:    9355a572b4e6218950a742097287942e104024585552e4356b2d144c5ce393bf
    Out:     0
    Script:  Jack
  Output 0:
    Value:  10
    Script: Ann
  Output 1:
    Value:  40
    Script: Jack
]
Prev: 0000da90f6f0987fb9b93a2ee8ca2b64be2a05e64c86f4f77961e8315f46dc9b
Hash: 0000968718061e70903ba823b5e6abc6f38c60d87680c006f92a109d1718a11d
Nonce: 6489
Pow: true

Timestamp: 1519799969
[Transaction 9355a572b4e6218950a742097287942e104024585552e4356b2d144c5ce393bf:
  Input 0:
    TXID:
    Out:     -1
    Script:  Decentralization make more people have more right
  Output 0:
    Value:  50
    Script: Jack
]
Prev:
Hash: 0000da90f6f0987fb9b93a2ee8ca2b64be2a05e64c86f4f77961e8315f46dc9b
Nonce: 50425
Pow: true

taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc send -from Jack -to Jack -amount 20
Mining block containing " [Transaction 3d474a4a8ba93a54eaad9964a666ae9a158dd49f4e82be2ba2dc44f526cebe6d:
  Input 0:
    TXID:    1256e1ca749b50a8d8bda37046b0ffceaf5f3bb2db019cff3141970fbd8bf405
    Out:     1
    Script:  Jack
  Output 0:
    Value:  20
    Script: Jack
  Output 1:
    Value:  20
    Script: Jack
] "


Success
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc printchain
Timestamp: 1519800050
[Transaction 3d474a4a8ba93a54eaad9964a666ae9a158dd49f4e82be2ba2dc44f526cebe6d:
  Input 0:
    TXID:    1256e1ca749b50a8d8bda37046b0ffceaf5f3bb2db019cff3141970fbd8bf405
    Out:     1
    Script:  Jack
  Output 0:
    Value:  20
    Script: Jack
  Output 1:
    Value:  20
    Script: Jack
]
Prev: 0000968718061e70903ba823b5e6abc6f38c60d87680c006f92a109d1718a11d
Hash: 0000807c3fd24bbed270df4fb9cd2920be6f2235ebd7949c3b05610c84585439
Nonce: 10754
Pow: true

Timestamp: 1519799993
[Transaction 1256e1ca749b50a8d8bda37046b0ffceaf5f3bb2db019cff3141970fbd8bf405:
  Input 0:
    TXID:    9355a572b4e6218950a742097287942e104024585552e4356b2d144c5ce393bf
    Out:     0
    Script:  Jack
  Output 0:
    Value:  10
    Script: Ann
  Output 1:
    Value:  40
    Script: Jack
]
Prev: 0000da90f6f0987fb9b93a2ee8ca2b64be2a05e64c86f4f77961e8315f46dc9b
Hash: 0000968718061e70903ba823b5e6abc6f38c60d87680c006f92a109d1718a11d
Nonce: 6489
Pow: true

Timestamp: 1519799969
[Transaction 9355a572b4e6218950a742097287942e104024585552e4356b2d144c5ce393bf:
  Input 0:
    TXID:
    Out:     -1
    Script:  Decentralization make more people have more right
  Output 0:
    Value:  50
    Script: Jack
]
Prev:
Hash: 0000da90f6f0987fb9b93a2ee8ca2b64be2a05e64c86f4f77961e8315f46dc9b
Nonce: 50425
Pow: true

taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc getbalance -address Jack
Balance of 'Jack': 80
taster@power:~/go/src/gobtc|master⚡
⇒  ./gobtc send -from Jack -to Jack -amount 80
Mining block containing " [Transaction d950b72b1a5fda9426ee7bfe716d5970eb3ca8801bf46035ba3f103808dd4bee:
  Input 0:
    TXID:    3d474a4a8ba93a54eaad9964a666ae9a158dd49f4e82be2ba2dc44f526cebe6d
    Out:     0
    Script:  Jack
  Input 1:
    TXID:    3d474a4a8ba93a54eaad9964a666ae9a158dd49f4e82be2ba2dc44f526cebe6d
    Out:     1
    Script:  Jack
  Input 2:
    TXID:    3d474a4a8ba93a54eaad9964a666ae9a158dd49f4e82be2ba2dc44f526cebe6d
    Out:     0
    Script:  Jack
  Input 3:
    TXID:    3d474a4a8ba93a54eaad9964a666ae9a158dd49f4e82be2ba2dc44f526cebe6d
    Out:     1
    Script:  Jack
  Output 0:
    Value:  80
    Script: Jack
] "


Success