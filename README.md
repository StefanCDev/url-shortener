url-shortener/
│── cmd/                      # Entry point for the app
│   └── main.go               # Main server file  
│
│── internal/                 # Business logic
│   ├── handlers/             # API request handlers  
│   │   ├── shorten.go        # Handles short URL creation  
│   │   ├── redirect.go       # Handles URL redirection  
│   │   └── health.go         # Health check endpoint  
│   │
│   ├── db/                   # Database interactions  
│   │   ├── dynamo.go         # DynamoDB helper functions  
│   │   ├── models.go         # Data structures (URL struct)  
│   │   └── config.go         # AWS DynamoDB configuration  
│   │
│   ├── utils/                # Utility functions  
│       ├── logger.go         # Logging helper (CloudWatch)  
│       ├── shortener.go      # Generates short URLs  
│       └── validator.go      # URL validation logic  
│
│── deployments/              # AWS Infrastructure  
│   ├── lambda/               # AWS Lambda function package  
│   ├── api-gateway.json      # API Gateway config  
│   └── dynamodb-schema.json  # DynamoDB table definition  
│
│── tests/                    # Unit tests  
│   ├── handlers_test.go      # Test API handlers  
│   ├── db_test.go            # Test database interactions  
│   └── utils_test.go         # Test helper functions  
│
│── .env                      # AWS credentials (use AWS IAM roles instead)  
│── go.mod                    # Go module dependencies  
│── README.md                 # Project documentation  
