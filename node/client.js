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
