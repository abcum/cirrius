export default {
    FIXTURES: [
        {
            id: "file:1",
            project: "project:1",
            path: "index.js",
            content: "var util = require('util');\n\nvar request = require('request');\n\nmodule.exports = function() {\n\n\trequest.post({\n\t\turl: 'http://requestb.in/',\n\t\tjson: {\n\t\t\ttime: new Date()\n\t\t}\n\t}, script.done);\n\n};",
            created_at: '2016-07-25T16:16:11.271Z',
            updated_at: '2016-07-25T16:16:24.695Z',
        },
        {
            id: "file:2",
            project: "project:1",
            path: "other.hbs",
        },
        {
            id: "file:3",
            project: "project:1",
            path: "app/main.js",
        },
        {
            id: "file:4",
            project: "project:1",
            path: "app/main.less",
        },
        {
            id: "file:5",
            project: "project:1",
            path: "app/templates/main.hbs",
        },
        {
            id: "file:6",
            project: "project:1",
            path: "controllers/project/main.js",
        },
    ]
};