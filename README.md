# Loco Backend Assessment - Transaction Service

## Overview
This project is a RESTful web service built using **Golang** and **Gin** framework. The service stores transaction data (in memory) and allows clients to retrieve transaction information, including calculating the sum of all transactions linked to a particular transaction via a parent-child relationship.

### Features
- **Create a transaction** with an optional parent.
- **Retrieve a transaction** by ID.
- **Retrieve all transactions of a given type**.
- **Calculate the total sum** of a transaction and all its linked (transitive) child transactions.

## API Endpoints

### 1. **Create a Transaction**
- **PUT** `/transactionservice/transaction/:transaction_id`
  
  Create a new transaction with an optional parent.
  
  **Request Body:**
  ```json
  {
    "amount": 5000,
    "type": "cars",
    "parent_id": 10 // Optional
  }
