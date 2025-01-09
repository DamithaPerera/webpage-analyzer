import React from "react";
import { Spinner } from "react-bootstrap";

const Loader = () => (
  <div className="d-flex justify-content-center align-items-center" style={{ minHeight: "100px" }}>
    <Spinner animation="border" variant="primary" />
  </div>
);

export default Loader;
