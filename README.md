# NFT HOMEWORK
## REQUIREMENTS
* Retrieve events from etherium blockchain by using etherscan.io
* Store events in a database PostgreSQL
* GraphQL Server for querying events
* No duplicate events read from the blockchain
## IMPLEMENTATION
1. Applied an API Key from etherscan.io for the project to access the etherscan.io.
2. Userd following APIs to serve the data:
    * [Transaction](https://api.etherscan.io/api?module=account&action=txlist&address=0x0&startblock=0&endblock=99999999&sort=asc&apikey=YourApiKeyToken)
    * [Internal Transaction](https://api.etherscan.io/api?module=account&action=txinternal&address=0x0&tag=latest&apikey=YourApiKeyToken)
    * [ERC20 Transaction](https://api.etherscan.io/api?module=account&action=tokentx&address=0x0&contractaddress=0x0&tag=latest&apikey=YourApiKeyToken)
    * [ERC721 Transaction](https://api.etherscan.io/api?module=account&action=tokennfttx&address=0x0&contractaddress=0x0&tag=latest&apikey=YourApiKeyToken)
    * [ERC1155 Transaction](https://api.etherscan.io/api?module=account&action=token1155tx&address=0x0&contractaddress=0x0&tag=latest&apikey=YourApiKeyToken)
    * [Current Block](https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=YourApiKeyToken)
3. Used PostgreSQL to store the events, and used GraphQL to query the events.
  * Postgres database schema is shown as below:
  ```
                                   Table "public.tokens"
      Column      |  Type  | Collation | Nullable |              Default        
       
------------------+--------+-----------+----------+-----------------------------
-------
 id               | bigint |           | not null | nextval('tokens_id_seq'::reg
class)
 name             | text   |           |          | 
 description      | text   |           |          | 
 contract_address | text   |           |          | 
 last_read_block  | bigint |           |          | 
 created_at       | bigint |           |          | 
Indexes:
    "tokens_pkey" PRIMARY KEY, btree (id)
    "tokens_contract_address_key" UNIQUE CONSTRAINT, btree (contract_address)
  ```

  ```
                                    Table "public.transactions"
       Column        |  Type  | Collation | Nullable |                 Default                
  
---------------------+--------+-----------+----------+----------------------------------------
--
 id                  | bigint |           | not null | nextval('transactions_id_seq'::regclass
)
 block_number        | text   |           |          | 
 time_stamp          | text   |           |          | 
 hash                | text   |           |          | 
 nonce               | text   |           |          | 
 block_hash          | text   |           |          | 
 transaction_index   | text   |           |          | 
 from                | text   |           |          | 
 to                  | text   |           |          | 
 value               | text   |           |          | 
 gas                 | text   |           |          | 
 gas_price           | text   |           |          | 
 is_error            | text   |           |          | 
 txreceipt_status    | text   |           |          | 
 input               | text   |           |          | 
 contract_address    | text   |           |          | 
 cumulative_gas_used | text   |           |          | 
 gas_used            | text   |           |          | 
 confirmations       | text   |           |          | 
 method_id           | text   |           |          | 
 function_name       | text   |           |          | 
Indexes:
    "transactions_pkey" PRIMARY KEY, btree (id)
    "transactions_hash_key" UNIQUE CONSTRAINT, btree (hash)
  ```


**NOTE: ** Schemas for the other 3 tables are not listed here.


4. Used Docker to run the server and database.

## RUN THE APPLICATION
1. Clone the repository to your local machine.
```
git clone https://github.com/tonyxu1/nfthomework
```
2. Run the following command to start the server:
```
$ cd nfthomework && docker-compose up --build -d
```
the command will start the GraphQL server and the database. The database will be created if it does not exist. The database `nft` and all tables  will be created if it does not exist.

3. After the server is started, open a Web browser and go to http://localhost:8080, and copy/past following qyery to the left panel and click on run button to see the result.


```
{
  getEvents(address: "testing", fromBlock: "1", toBlock: "2") {
    currentBlock,
    transactions {
      blockNumber
      timeStamp
      hash
      nonce
      blockHash
      transactionIndex
      from
      to
      value
      gas
      gasPrice
      isError
      txreceipt_status
      input
      contractAddress
      cumulativeGasUsed
      gasUsed
      confirmations
      methodId
      functionName
    },
    trxInternal {
      blockNumber
      timeStamp
      hash
      from
      to
      value
      contractAddress
      input
      type
      gas
      gasUsed
      traceId
      isError
      errCode
    },
    trxERC20 {
      blockNumber
      timeStamp
      hash
      nonce
      blockHash
      from
      contractAddress
      to
      value
      tokenName
      tokenSymbol
      tokenDecimal
      transactionIndex
      gas
      gasPrice
      gasUsed
      cumulativeGasUsed
      input
      confirmations
    },
    trxERC721 {
      blockNumber
      timeStamp
      hash
      nonce
      blockHash
      from
      contractAddress
      to
      tokenID
      tokenName
      tokenSymbol
      tokenDecimal
      transactionIndex
      gas
      gasPrice
      gasUsed
      cumulativeGasUsed
      input
      confirmations
    },
    trxERC1155 {
      blockNumber
      timeStamp
      hash
      nonce
      blockHash
      transactionIndex
      gas
      gasPrice
      gasUsed
      cumulativeGasUsed
      input
      contractAddress
      from
      to
      tokenID
      tokenValue
      tokenName
      tokenSymbol
      confirmations
    }
  }
}
```

**Note:** The parameters are just fake data, all parameters will be read from the database.

## TODOS:
1. Unit testing
2. Multiple contract addresses support
3. Query the events from Database first if block number is smaller than the last read block number; otherwise, query the events from Etherscan.
