const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = './tasks.proto';
const SERVER_ADDRESS = 'localhost:50051';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const { tasks: tasksDescriptor } = protoDescriptor;

const tasks = [];
const addTask = (call, callback) => {
  tasks.push(call.request);
  console.log('added 1 task');
  callback(null);
}
const listTasks = (call) => {
  tasks.forEach((task) => call.write(task));
  console.log(`listed ${tasks.length} tasks`);
  call.end();
}

const server = new grpc.Server();
server.addService(tasksDescriptor.Tasks.service, {
  addTask,
  listTasks,
});
server.bindAsync(SERVER_ADDRESS, grpc.ServerCredentials.createInsecure(), () => {
  server.start();
  console.log(`server running at ${SERVER_ADDRESS}`);
});
