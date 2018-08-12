import { axios } from 'axios';

export class UserService {
    static async createUser(newUser) {

        console.log('[LOG] all', newUser);
        return new Promise((resolve, reject) => {
            if (!newUser) {
                reject('Empty user');
                return;
            }
            setTimeout(() => {
                resolve(newUser);
            }, 1000);
        })
        // return axios
        //     .post('', newUser);
    }
}
