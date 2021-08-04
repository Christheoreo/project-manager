import type { ILoginResponse } from './../dtos/ILoginResponseDto';
import { Service } from './Service';

export class AuthenticationService extends Service {
  constructor() {
    super('auth');
  }

  public async login(email: string, password: string): Promise<ILoginResponse> {
    const { data } = await this.instance.post('login', { email, password });
    return data;
  }
}
