<script lang="ts">
import router from 'page';

import { Container, Row, Col, Form, FormGroup, FormText, Input, Label, Button, Alert } from 'sveltestrap';
import type { IStandardResponse } from '../dtos/IStandardResponse';
import type { IError } from '../interfaces/IError';
import { AuthenticationService } from '../services/AuthenticationService';
import { IsLoggedInStore } from '../stores/IsLoggedInStore';
import { Validation } from '../utils/Validation';
import Loader from './../components/Loader.svelte';
const authenticationService = new AuthenticationService();

let email: string = 'chris.finlow@gmail.com';
let password: string = '12345678';

let showErrorModal: boolean = false;
let errorMessage: string = '';

let showSpinner: boolean = false;
let isLoggingIn: boolean = false;

const validateUserInput: () => IError = (): IError => {
  let res: IError = { error: true, message: 'NA' };
  if (!Validation.isEmailValid(email)) {
    res.message = 'Email is invalid.';
  } else if (!Validation.isPasswordValid(password)) {
    res.message = 'Password needs to be between 8 and 25 characters';
  } else {
    res.error = false;
  }

  return res;
};

const login = async () => {
  if (isLoggingIn) return;
  isLoggingIn = true;
  showErrorModal = false;
  showSpinner = false;
  let validationResponse = validateUserInput();

  if (validationResponse.error) {
    //
    errorMessage = validationResponse.message;
    showErrorModal = true;
    isLoggingIn = false;
    return;
  }
  showSpinner = true;

  try {
    const loginResponse = await authenticationService.login(email, password);
    window.localStorage.setItem('token', loginResponse.accessToken);
    IsLoggedInStore.set(true);
    router.replace('/');
  } catch (error) {
    if (error.response) {
      const loginError: IStandardResponse = error.response.data;
      if (loginError.messages.length == 1) {
        errorMessage = `${loginError.messages[0]} :(`;
      } else {
        errorMessage = loginError.messages.join(' || ');
      }
    } else {
      errorMessage = 'Unknown error :(';
    }
    showErrorModal = true;
  } finally {
    showSpinner = false;
  }
  //
  isLoggingIn = false;
};

const handleEnterKey = async (e: KeyboardEvent) => {
  if (e.code === 'Enter') {
    await login();
  }
};
</script>

<svelte:head>
  <title>Project Manager - Login</title>
</svelte:head>

<Container class="h-100">
  <Row class="h-100 align-items-center justify-content-center">
    <Col xs="{12}" md="{6}">
      <h1>Project Manager</h1>
      <Form>
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
          <Button color="dark" block="{true}" type="button" on:click="{login}">Login</Button>
        </FormGroup>
        <Loader show="{showSpinner}" />
        <Alert color="warning" isOpen="{showErrorModal}" dismissible>{errorMessage}</Alert>
        <hr />
        <FormGroup>
          <p class="text-muted">Need an account? <a class="link-dark" href="/register">Register</a></p>
        </FormGroup>
      </Form>
    </Col>
  </Row>
</Container>
