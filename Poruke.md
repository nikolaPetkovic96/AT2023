# Poruke
### <u>**Consumer**</u>
#### Recieves:
- **SuccessfulTransaction** - Ends transaction
#### Sends:
- **BuyProduct** - Starts transaction


### <u>DistributorActor</u>
#### Recieves:
- **BuyProduct** - Starts process
- **MissingProduct** - Listens for missing products in storages and organizes to get product to that storage
- **ProductAvailable** - Gets availablity of products in storages
- **ProductToDistributor** - Recieves product and sends **SuccessfulTransaction** to consumer
- **ProductCreated** - Gets newly created product and distributes it where its needed
#### Sends:
- **CheckProduct** - Check for a product
- **GetProduct** - Gets product (removes it from the storage)
- **SuccessfulTransaction** - response to consumer on successful transaction
- **ProductToRetailer** - Adds product to the storage
- **NewProduct** - when there is not enough products, gets new one from provider


### <u>StateManagmentActor</u>
#### Recieves:
- **RefreshState** - Executes script that gets and saves retailers inventory and self pings with delay 
- **InventoryResponse** - Inventory data got from retailer, saves them locally
#### Sends:
- **RefreshState** - Self ping with a delay
- **GetInventory** - Sends to retailer actor to get all their inventory

### <u>RetailerActor</u>
#### Recieves:
- **GetInventory** - Gets all data in the storage and sends **InventoryResponse** with data
- **CheckProduct** - Checks for specific product and quantity 
- **GetProduct** - Send for specific product and quantity to **DistributorActor**
- **CheckInventoryStatus** - Checks inventory for missing products, sends **  
- **ProductToRetailer** - Gets product and saves to database
#### Sends:
- **InventoryResponse** - Message with inventory data
- **CheckInventoryStatus** - Self ping with a delay
- **MissingProduct** - Send info about missing product to **DistributorActor**
- **ProductAvailable** - Response to check product
- **ProductToDistributor** - Sends product to distributor

### <u>SuplierActor</u>
#### Recieves:
- **NewProduct** - Creates a new product
#### Sends:
- **ProductCreated** - Sends newly created products back to distributor
