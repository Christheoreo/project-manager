<script lang="ts">
import { Container, Row, Col, Form, FormGroup, FormText, Input, Label, Button } from 'sveltestrap';

let email: string = '';
let firstName: string = '';
let lastName: string = '';
let password: string = '';
let passwordConfirm: string = '';

const register = async () => {
  alert(email);
  alert(password);
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
        <FormGroup>
          <Button color="dark" block="{true}" type="button" on:click="{register}">Login</Button>
        </FormGroup>
        <hr />
        <FormGroup>
          <p class="text-muted">Have an account? <a class="link-dark" href="/login">Login</a></p>
        </FormGroup>
      </Form>
    </Col>
  </Row>
</Container>
