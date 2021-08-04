export interface IRoute {
  path: string;
  component: any;
  protected?: boolean;
  redirectIfLoggedIn?: boolean;
  children?: IRoute[];
}
