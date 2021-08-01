<script lang="ts">
import { Container, Row, Col, Form, FormGroup, FormText, Input, Label, Button, Alert } from 'sveltestrap';
import { Validation } from '../utils/Validation';

let email: string = '';
let password: string = '';

let showErrorModal: boolean = false;
let errorMessage: string = '';

const validateUserInput: () => { error: boolean; message: string } = (): { error: boolean; message: string } => {
  let res: { error: boolean; message: string } = { error: true, message: 'NA' };
  if (!Validation.isEmailValid(email)) {
    res.message = 'Email is invalid.';
  } else if (!Validation.isPasswordValid(password)) {
    res.message = 'Password needs to be between 8 and 25 characters';
  }

  return res;
};

/**
 * Carry on here. 
 * Add a loading spinner as well!
 */
const login = async () => {
  showErrorModal = false;
  let validationResponse = validateUserInput();

  if (validationResponse.error) {
    //
    errorMessage = validationResponse.message;
    showErrorModal = true;
  }
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
        <Alert color="warning" isOpen="{showErrorModal}" dismissible>{errorMessage}</Alert>
        <hr />
        <FormGroup>
          <p class="text-muted">Need an account? <a class="link-dark" href="/register">Register</a></p>
        </FormGroup>
      </Form>
    </Col>
  </Row>
</Container>
