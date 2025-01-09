import React from "react";
import { Navbar, Container } from "react-bootstrap";

const Header = () => (
  <Navbar bg="dark" variant="dark" expand="lg" sticky="top">
    <Container>
      <Navbar.Brand href="/">🌐 Web Page Analyzer</Navbar.Brand>
    </Container>
  </Navbar>
);

export default Header;
