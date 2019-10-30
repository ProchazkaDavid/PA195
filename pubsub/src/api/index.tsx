import Redis from 'redis';
import config from './config.json';

const connect = () => {
    const client = Redis.createClient(config);
    client.on('error', err => {
        console.log(`There has been a murder! - ${err}`);
    });

    client.set('key', 'value', Redis.print);
    client.get('key', value => {
        console.log(`this is the value: ${value}`);
    });
};