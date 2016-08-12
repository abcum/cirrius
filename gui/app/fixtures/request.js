export default {
	FIXTURES: [
		{
			id: "request:1",
			project: "project:1",
			status: 200,
			url: "http://www.apple.com/",
			duration: 1283719281,
			ip: "27.0.0.1",
			method: "GET",
			size: 123456,
			time: "2016-07-25T16:16:11.271Z",
			logs: [
				"log:1",
				"log:2",
			],
		},
		{
			id: "request:2",
			project: "project:1",
			status: 200,
			url: "http://www.apple.com/",
			duration: 1283719281,
			ip: "27.0.0.1",
			method: "GET",
			size: 123456,
			time: '2016-07-25T16:16:11.271Z',
			logs: [
				"log:3",
				"log:4",
			],
		},
	]
};
