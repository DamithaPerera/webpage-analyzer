import React from "react";
import { Spinner } from "react-bootstrap";

const Loader = () => (
  <div className="d-flex justify-content-center align-items-center">
    <Spinner animation="border" />
  </div>
);

export default Loader;
