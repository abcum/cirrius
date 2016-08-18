export default {
	FIXTURES: [
		{
			id: "log:1",
			type: "info",
			fold: null,
			file: null,
			line: null,
			char: null,
			args:[
				"Started script",
				"2016-07-25T16:16:11.271Z"
			],
		},
		{
			id: "log:2",
			type: "log",
			fold: "app/",
			file: "index.js",
			line: 15,
			char: 1,
			args:[
				"Submitting sms to AWS SQS"
			],
		},
		{
			id: "log:3",
			type: "warn",
			fold: "app/",
			file: "index.js",
			line: 17,
			char: 1,
			args:[
				"Callback function deprecated"
			],
		},
		{
			id: "log:4",
			type: "error",
			fold: "app/",
			file: "index.js",
			line: 18,
			char: 2,
			args:[
				"Module 'ember' not found"
			],
		},
		{
			id: "log:5",
			type: "info",
			fold: null,
			file: null,
			line: null,
			char: null,
			args:[
				"Finished script",
				"2016-07-25T16:16:11.521Z"
			],
		},
		{
			id: "log:6",
			type: "time",
			fold: null,
			file: null,
			line: null,
			char: null,
			args:[
				"Script duration",
				250*1000000
			],
		},
	]
};
