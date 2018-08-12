import { axios } from 'axios';

export class LoginService {
    static async login(creds) {
        return new Promise((resolve, reject) => {
            console.log('[LOG] login creds', creds);
            if (creds.password === 'nok') {
                throw new Error('Incorrect password');
            }
            resolve('my.new.token');
            // axios
            // .post('/login', {
            //     username: creds.username,
            //     password: creds.password
            // })
            // .then((response) => {
            //     resolve(reponse);
            //     console.log(response);
            // })
            // .catch((error) => {
            //     reject(error);
            //     console.log(error);
            // });
        })
    }
}