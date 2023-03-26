const grpc = require('@grpc/grpc-js');
const createTasksDescriptor = require('./descriptor');
const { PROTO_PATH, SERVER_ADDRESS } = require('./config');

const tasksDescriptor = createTasksDescriptor(PROTO_PATH);
const stub = new tasksDescriptor.Tasks(SERVER_ADDRESS, grpc.credentials.createInsecure());
stub.addTask({
  title: 'first task',
  description: 'hello world',
  status: 0
}, (err) => {
  if (err) {
    console.log(`error while adding task: ${err}`);
    return;
  }
  
  const call = stub.listTasks();
  let tasksCount = 0;
  call.on('error', (err) => {
    console.error(err);
    process.exit(1);
  });
  call.on('data', (data) => {
    console.log(`got back a task: ${JSON.stringify(data)}`);
    tasksCount += 1;
  });
  call.on('end', () => {
    console.log(`received ${tasksCount} tasks`);
  });
});
