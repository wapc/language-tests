const fs = require("fs");
const path = require("path");
const url = require("url");
const https = require("https");
const HttpsProxyAgent = require("https-proxy-agent");
const yaml = require("js-yaml");
const prettier = require("prettier");
const child_process = require("child_process");
const { parse, Context, Writer } = require("widl-codegen/widl");

const args = process.argv.slice(2);
if (args.length < 1) {
  console.log("usage: widl-codegen <configfile>");
  process.exit(1);
}

try {
  generate(args[0]);
} catch (err) {
  console.log(err);
}

function generate(configFile) {
  let configContents = fs.readFileSync(configFile, "utf8");
  let config = yaml.safeLoad(configContents);
  let parentDir = config.parentDir || "";
  parentDir = parentDir.trim();
  if (parentDir.length == 0) {
    parentDir = "./";
  } else if (!parentDir.endsWith(path.sep)) {
    parentDir += path.sep;
  }

  load(config.schema, function (err, data) {
    if (err) {
      console.error(err);
      return;
    }

    const schemaContents = data;
    const doc = parse(schemaContents);

    for (var entry of Object.entries(config.generates)) {
      const filename = entry[0];
      const fullPath = parentDir + filename;
      const generate = entry[1];

      if (generate.ifNotExists && fs.existsSync(fullPath)) {
        return;
      }

      const writer = new Writer();
      const context = new Context(generate.config);
      const pkg = require(generate.package);
      const visitorClass = pkg[generate.visitorClass];
      const visitor = new visitorClass(writer);
      doc.accept(context, visitor);
      let source = writer.string();
      const ext = path.extname(fullPath).toLowerCase();
      switch (ext) {
        case ".ts":
          source = formatAssemblyScript(source);
          break;
      }
      fs.writeFileSync(fullPath, source);
      switch (ext) {
        case ".go":
          formatGolang(filename, parentDir);
          break;
        case ".rs":
          formatRust(filename, parentDir);
          break;
      }
    }
  });
}

async function load(endpoint, callback) {
  if (endpoint.startsWith("http://") || endpoint.startsWith("https://")) {
    var proxy = process.env.http_proxy;
    var options = url.parse(endpoint);
    if (proxy) {
      const agent = new HttpsProxyAgent(proxy);
      options.agent = agent;
    }
    const req = https.request(options, (res) => {
      let response = "";
      res.on("data", (d) => {
        response += d.toString();
      });
      res.on("end", (d) => {
        if (res.statusCode / 100 == 2) {
          callback(null, response);
        }
      });
    });
    req.on("error", (e) => {
      callback(e);
    });
    req.end();
  } else {
    fs.readFile(endpoint, "utf8", callback);
  }
}

function formatAssemblyScript(source) {
  try {
    source = prettier.format(source, {
      semi: true,
      parser: "typescript",
    });
  } catch (err) {}
  return source;
}

function formatGolang(filename, cwd) {
  child_process.execSync("go fmt " + filename, { cwd: cwd });
}

function formatRust(filename, cwd) {
  child_process.execSync("rustfmt " + filename, { cwd: cwd });
}
