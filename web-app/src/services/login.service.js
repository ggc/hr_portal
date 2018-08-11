import { resolve } from "url";

export class LoginService {
    static async login(creds) {
        return new Promise((resolve, reject) => {
            console.log('[LOG] login creds', creds);
            if (creds.password === 'nok') {
                throw new Error('Incorrect password');
                reject('Incorrect pass, man');
            }
            setTimeout(() => {
                resolve('myToken');
            }, 1000);
        })
    }
}