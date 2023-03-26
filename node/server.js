const grpc = require('@grpc/grpc-js');
const createTasksDescriptor = require('./descriptor');
const { PROTO_PATH, SERVER_ADDRESS } = require('./config');

const tasksDescriptor = createTasksDescriptor(PROTO_PATH);

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
