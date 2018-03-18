# 一个简单的加密货币理论及实现

## 前言
我参照的博客:<https://jeiwan.cc/tags/blockchain/>

## 实现了的特性

* 区块链的实现

    * 链式存储结构
    * 工作量证明
    * 数据库存储
    * 浏览和添加区块

* 交易的实现

    * 交易结构
    * coinbase 交易
    * 查询余额
    * 普通交易

* 钱包实现

    * 钱包生成
    * 交易的签名和校验

## 未实现的特性

* 对于上述实现的算法及数据结构优化
* 网络特性

---

## 区块链

### 原理

#### Question: 如何在去信任的条件下保证一个电子数据不被篡改?

回答 1: 把文件加上 hash 

* 原理:

        目前sha256碰撞在计算上是不可行的  
* 缺点

回答 2: 分布式存储,加链式结构

* 原理:
    * 链式结构

        |timestamp | data | prevhash
        |- | :-: | -:
        |0 | Gryffindor| nil
        |1 | Ann | hash0
        |2 | Slytherin | hash1
        |3 | geaiu| hash2
        |4 | gagharh | hash3
        |..|
        |n | egaga | hashn-1
        now

        修改数据必然导致接下来的整个链不同,分布式存储保证链不会被丢失

* 缺点

回答 3: 分布式存储, 加工作量验证,加最长链原则

* 原理
    * 结构
        |timestamp | data | prevhash|nonce
        |- | :-: | -:|-
        |0 | Gryffindor| nil|nonce0
        |1 | Ann | hash0|nonce1
        |2 | Slytherin | hash1|nonce2
        |3 | geaiu| hash2|nonce3
        |4 | gagharh | hash3|nonce4
        |...|
        |n | egaga | hashn-1|noncen
        |now
    * 工作量证明使 区块的hash必须满足一定条件,使添加区块变得困难
    * 若不能超过50%的算力,则不能构造出比现有链更长的链
    * 由此可使链链上的数据无法篡改
* 缺点
    必须防范50%攻击

回答 4: 直接离线存储

### 实现

### 小评

---

## 交易

### 原理

#### 目标:在区块链上实现交易,保证不能被双花

#### 实现:将交易视为输入和输出

~~~go
type Transaction struct {
    ID   []byte
    Vin  []TXInput
    Vout []TXOutput
}

type TXOutput struct {
    Value        int
    ScriptPubKey string //目前为接收者用户钱包地址
}

type TXInput struct {
    Txid      []byte
    Vout      int
    ScriptSig string //目前位在发送者用户钱包地址
}
~~~

![transaction](transactions-diagram.png)

如此结构可以保证在不修改数据的情况下进行交易

#### 查询余额

#### 交易

#### 双花的可能性

---

## 钱包

### 钱包的生成原理

~~~txt
助记词
  |
  V
私钥1, 私钥2, 私钥3,...
  |     |      |
  V     V      V
公钥1, 公钥2, 公钥3,...
  |     |      |
  V     V      V
地址1, 地址2, 地址3...
~~~

地址 = 地址版本+公钥hash+冗余校验

### 签名和校验

采用ecdsa签名算法
![sign](signing-scheme.png)

* 签名保证无法盗用别人的钱包花钱
* 也保证花的钱无法抵赖
* 校验失败的交易不加入区块
* 交易的结构保证不存在公钥欺骗

~~~go
type TxOutput struct {
    Value      int
    PubKeyHash []byte
}

type TxInput struct {
    Txid      []byte //引用TxOut的交易id
    Vout      int    // 引用TxOut在交易中的输出序号
    Signature []byte
    PubKey    []byte
}
~~~

---

## 提问
