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
import { onMount } from 'svelte';
import { UsersService } from '../services/UsersService';
const usersService = new UsersService();
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
  await getUserInfo();
});
</script>

<Navbar color="light" light expand="md">
  <NavbarBrand href="/">Project Manager</NavbarBrand>
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
          <DropdownItem href="/logout">Logout</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </Nav>
  </Collapse>
</Navbar>
<slot />
