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
git clone 
```