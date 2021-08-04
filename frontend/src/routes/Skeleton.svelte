<script lang="ts">
import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  Dropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
} from 'sveltestrap';
import { navigate, link } from 'svelte-routing';
import { onMount } from 'svelte';
import { UsersService } from '../services/UsersService';
const usersService = new UsersService();
const login = () => navigate('/login', { replace: true });
let isOpen = false;
export let name: string = 'Options';

const handleUpdate = (event) => {
  isOpen = event.detail.isOpen;
};

const getUserInfo = async () => {
  try {
    const user = await usersService.getMe();
    name = `${user.firstName} ${user.lastName}`;
  } catch (error) {
    //
  }
};

onMount(async () => {
  console.log('here');
  await getUserInfo();
});
</script>

<Navbar color="light" light expand="md">
  <NavbarBrand href="/">sveltestrap</NavbarBrand>
  <NavbarToggler on:click="{() => (isOpen = !isOpen)}" />
  <Collapse isOpen="{isOpen}" navbar expand="md" on:update="{handleUpdate}">
    <Nav class="ms-auto" navbar>
      <NavItem>
        <NavLink href="#components/">Components</NavLink>
      </NavItem>
      <NavItem>
        <NavLink href="https://github.com/bestguy/sveltestrap">GitHub</NavLink>
      </NavItem>
      <Dropdown nav inNavbar>
        <DropdownToggle nav caret>{name}</DropdownToggle>
        <DropdownMenu end>
          <DropdownItem><a href="/logout" use:link>Logout</a></DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </Nav>
  </Collapse>
</Navbar>

<slot />
