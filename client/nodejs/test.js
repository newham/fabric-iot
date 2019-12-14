// 测试用客户端，测试系统接口性能

var express = require('express')
var app = express()

const { invoke } = require("./invoke")

// respond with "hello world" when a GET request is made to the homepage
app.get('/fabric-iot/test', async function (req, res) {
    const cc_name = req.query.cc_name;
    const f_name = req.query.f_name;
    const args = req.query.args;
    // console.log("invoke",cc_name,f_name,args);
    if (cc_name == undefined || f_name == undefined || args == undefined) {
        res.status(400).send({ status: 400, msg: 'bad params' });
        return
    }
    const r = await invoke(cc_name, f_name, args.split("|"));
    res.status(r.status).send(r);
})

const port = 8001
app.listen(port, () => console.log(`Test invoke on port ${port}!`))