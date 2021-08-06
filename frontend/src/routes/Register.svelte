<script lang="ts">
import {
  Container,
  Row,
  Col,
  Form,
  FormGroup,
  FormText,
  Input,
  Label,
  Button,
  Alert,
  ListGroup,
  ListGroupItem,
} from 'sveltestrap';
import router from 'page';
import type { INewUserDto } from '../dtos/INewUserDto';
import type { IStandardResponse } from '../dtos/IStandardResponse';
import type { IUserResponse } from '../dtos/IUserResponse';
import type { IError } from '../interfaces/IError';
import type { IValidationError } from '../interfaces/IValidationError';
import { AuthenticationService } from '../services/AuthenticationService';
import { UsersService } from '../services/UsersService';
import { IsLoggedInStore } from '../stores/IsLoggedInStore';
import { Validation } from '../utils/Validation';
import Loader from './../components/Loader.svelte';
const usersService = new UsersService();
const authenticationService = new AuthenticationService();
let email: string = '';
let firstName: string = '';
let lastName: string = '';
let password: string = '';
let passwordConfirm: string = '';

let registrationInProgress: boolean = false;
let showErrorModal: boolean = false;
let errorMessages: string[] = [];

let showSpinner: boolean = false;

const validate = (): IValidationError => {
  let res: IValidationError = { error: true, messages: [] };
  if (firstName.length < 3) {
    res.messages.push('First name needs to be between 3 and 30 characters');
  }
  if (lastName.length < 3) {
    res.messages.push('Last name needs to be between 3 and 30 characters');
  }
  if (!Validation.isEmailValid(email)) {
    res.messages.push('Email is invalid.');
  }
  if (!Validation.isPasswordValid(password)) {
    res.messages.push('Password needs to be between 8 and 25 characters');
  }
  if (password !== passwordConfirm) {
    res.messages.push('Passwords do not match.');
  }

  console.log(res)
  res.error = res.messages.length > 0;

  return res;
};

const register = async () => {
  if (registrationInProgress) return;
  registrationInProgress = true;

  const validationResponse = validate();

  if (validationResponse.error) {
    errorMessages = validationResponse.messages;
    showErrorModal = true;
    registrationInProgress = false;
    return;
  }
  showSpinner = true;

  let user: IUserResponse;
  try {
    let newUserDto: INewUserDto = {
      email,
      firstName,
      lastName,
      password,
      passwordConfirm,
    };
    user = await usersService.register(newUserDto);
  } catch (error) {
    if (error.response) {
      const registrationError: IStandardResponse = error.response.data;
      if (registrationError.messages.length == 1) {
        errorMessages = [`${registrationError.messages[0]} :(`];
      } else {
        errorMessages = registrationError.messages;
      }
    } else {
      errorMessages = ['Unknown error :('];
    }
    showErrorModal = true;
    showSpinner = false;
    registrationInProgress = false;
    return;
  }

  // Now login the user!

  try {
    const loginResponse = await authenticationService.login(user.email, password);
    window.localStorage.setItem('token', loginResponse.accessToken);
    IsLoggedInStore.set(true);
    router.replace('/');
  } catch (error) {
    if (error.response) {
      const loginError: IStandardResponse = error.response.data;
      if (loginError.messages.length == 1) {
        errorMessages = [`${loginError.messages[0]} :(`];
      } else {
        errorMessages = loginError.messages;
      }
    } else {
      errorMessages = ['Account was created, but unable to login :('];
    }
    showErrorModal = true;
    showSpinner = false;
    registrationInProgress = false;
  }
};

const handleEnterKey = async (e: KeyboardEvent) => {
  if (e.code === 'Enter') {
    await register();
  }
};
</script>

<svelte:head>
  <title>Project Manager - Register</title>
</svelte:head>

<Container class="h-100">
  <Row class="h-100 align-items-center justify-content-center">
    <Col xs="{12}" md="{6}">
      <h1>Project Manager</h1>
      <Form>
        <FormGroup>
          <Label for="firstName">First Name</Label>
          <Input
            type="text"
            name="firstName"
            id="firstName"
            placeholder="Joe"
            bind:value="{firstName}"
            on:keyup="{handleEnterKey}" />
        </FormGroup>
        <FormGroup>
          <Label for="lastName">Last Name</Label>
          <Input
            type="text"
            name="lastName"
            id="lastName"
            placeholder="Bloggs"
            bind:value="{lastName}"
            on:keyup="{handleEnterKey}" />
        </FormGroup>
        <FormGroup>
          <Label for="email">Email</Label>
          <Input
            type="email"
            name="email"
            id="email"
            placeholder="email@example.com"
            bind:value="{email}"
            on:keyup="{handleEnterKey}" />
        </FormGroup>
        <FormGroup>
          <Label for="password">Password</Label>
          <Input
            type="password"
            name="password"
            id="password"
            placeholder="********"
            bind:value="{password}"
            on:keyup="{handleEnterKey}" />
        </FormGroup>
        <FormGroup>
          <Label for="passwordConfirm">Password Confirm</Label>
          <Input
            type="password"
            name="passwordConfirm"
            id="passwordConfirm"
            placeholder="********"
            bind:value="{passwordConfirm}"
            on:keyup="{handleEnterKey}" />
        </FormGroup>
        <Loader show="{showSpinner}" />
        <Alert color="warning" isOpen="{showErrorModal}" dismissible>
          <ListGroup flush>
            {#each errorMessages as errorMessage}
              <ListGroupItem color="warning">{errorMessage}</ListGroupItem>
            {/each}
          </ListGroup>
        </Alert>
        <hr />
        <FormGroup>
          <Button color="dark" block="{true}" type="button" on:click="{register}">Register</Button>
        </FormGroup>
        <hr />
        <FormGroup>
          <p class="text-muted">Have an account? <a class="link-dark" href="/login">Login</a></p>
        </FormGroup>
      </Form>
    </Col>
  </Row>
</Container>
