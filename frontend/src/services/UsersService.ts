import type { INewUserDto } from '../dtos/INewUserDto';
import type { IUserResponse } from './../dtos/IUserResponse';
import { Service } from './Service';

export class UsersService extends Service {
  constructor() {
    super('users');
  }

  public async getMe(): Promise<IUserResponse> {
    const { data } = await this.instance.get('me');
    return data;
  }
  public async register(newUser: INewUserDto): Promise<IUserResponse> {
    const { data } = await this.instance.post('register', newUser);
    return data;
  }
}
