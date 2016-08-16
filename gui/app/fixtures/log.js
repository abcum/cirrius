export default {
    FIXTURES: [
        {
            id: "log:1",
            type: "time",
            fold: "app/",
            file: "index.js",
            line: 12,
            char: 1,
            args:[
            	"My function",
            	250,
            ],
        },
        {
            id: "log:2",
            type: "log",
            fold: "app/",
            file: "index.js",
            line: "15",
            char: "1",
            args:[],
        },
        {
            id: "log:3",
            type: "log",
            fold: "app/",
            file: "index.js",
            line: "17",
            char: "1",
            args:[],
        },
        {
            id: "log:4",
            type: "log",
            fold: "app/",
            file: "index.js",
            line: "18",
            char: "2",
            args:[],
        },
    ]
};
