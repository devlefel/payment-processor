# Payment Processor

A Simple payment processor, that tries to follow pci rules <https://pt.pcisecuritystandards.org/pci_security/>

## How to work

- ```git clone $project```
- ```cd $project```
- ```docker-compose -f "docker-compose.yml" up -d --build```
- The tests will run and the containers will be up after all of them pass

## Endpoints
- Login Endpoint:
    - ```POST http://localhost:8081/user/login```
    - JSON BODY:
        - ```json
          {
            "username": "admin",
            "password": "@#$RF@!718"
          }
          ```
    - The response Will contain a array of strings, which each one represents a token, we need one of them for the next endpoint

- Process Endpoint
    - This Endpoint is the endpoint to process a purchase
    - ```POST http://localhost:8082/process/payment```
    - JSON BODY:
        - ```json {
                  	"token" : "eb2cd9bf2054de63b62330b3ae319e517f195afcc0ed19e984910f833d7f95a2",
                  	"card" : {
                  	"open" : {
                  		"name": "Felipe Gomes",
                  		"flag": "Visa",
                  		"date": "09/23"
                  	}
                  	},
                  	"process": {
                  		"totalvalue": 1000.00,
                  		"items": [
                  			{
                  				"name": "Geladeira",
                  				"value": 1000.00
                  			}
                  		],
                  		"installments": 10,
                  		"seller": {
                  			"name": "Magazine Luiza",
                  			"cnpj": "39.890.918/0001-99",
                  			"address": {
                  				"street": "Rua Jeronimo da Veiga",
                  				"number": 185,
                  				"zipcode": "04812-190"
                  			}
                  		}
                  	},
                  	"acquirer_id" : 2
                  }
          
- It will return the following JSON RESPONSE {
```json
{
  "success": true,
  "errors": {
    "Validation": null,
    "Internal": null
  }
}
```

## Contact
Email: felipe.pbgomes@gmail.com