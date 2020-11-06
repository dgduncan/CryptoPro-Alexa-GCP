var async = require("async");
var aws = require("aws-sdk");
var converter = require('number-to-words');
var dynamodb = new aws.DynamoDB();
const uuidv4 = require('uuid/v4')

var currencies = ["Bitcoin", "Ethereum", "Bitcoin Cash"];
var mainText = "";

exports.handler = (event, context, callback) => {
	
	var queue = async.queue(function (task, queueCallback) {	
		getCoin(task.coin, queueCallback)
	});
		
    for (var index = 0; index <= 2; index++) {
		queue.push({coin : currencies[index]}, function(error){
			if(error)
				console.log(error)
			});
	}
	
	queue.drain = function() {
		buildAndSendResponse(callback)
	}
}
        
function getCoin(coin, queueCallback) {
    var params = {
        TableName: "CryptoPro",
        Key: {
            "name": {
                S: coin.toLowerCase()
            },
        }
    };
    
    dynamodb.getItem(params, function(err, data) {
        if (err)
            console.log(err, err.stack); // an error occurred
        else {
            buildMainText(data, queueCallback)
        }
    });
}

function buildMainText(data, queueCallback) {
    var priceChangeDirection = "down";
    
    if (data.Item.percent_change_day.N > 0) 
        priceChangeDirection = "up";
        
    var coinText = "The price of " + data.Item.name.S + " is " + priceChangeDirection + " by " + converter.toWords(data.Item.percent_change_day.N) + " percent to a price of " + converter.toWords(data.Item.price.N) + " dollars. ";

    mainText += coinText
    queueCallback()
}
function buildAndSendResponse(callback) {
    callback(null, {
        statusCode: '200',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(
            {
                "uid": uuidv4(),
                "updateDate": new Date(),
                "titleText": "Today's most recent Crypto pricing Update",
                "mainText": mainText})
            });
	mainText = "";
}