import React, { useState } from "react";
import { Container, Form, Button, Alert, Card, Row, Col } from "react-bootstrap";
import { isValidURL } from "../utils/validation";
import { analyzeURL } from "../api/apiService";
import Loader from "../components/Loader";

const HomePage = () => {
  const [url, setURL] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState(null);

  const handleAnalyze = async (e) => {
    e.preventDefault();
    setError("");
    setResult(null);

    if (!isValidURL(url)) {
      setError("Invalid URL format. Please enter a valid URL.");
      return;
    }

    try {
      setLoading(true);
      const data = await analyzeURL(url);
      setResult(data);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container className="py-5">
      <Row>
        <Col md={8} className="mx-auto">
          <Card className="shadow-sm">
            <Card.Body>
              <h1 className="text-center mb-4">Web Page Analyzer</h1>
              <p className="text-muted text-center mb-4">
                Analyze the metadata, headings, links, and more of any webpage.
              </p>
              <Form onSubmit={handleAnalyze} className="d-flex flex-column gap-3">
                <Form.Group controlId="urlInput" className="mb-3">
                  <Form.Label className="fw-bold">Enter URL</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="https://example.com"
                    value={url}
                    onChange={(e) => setURL(e.target.value)}
                    isInvalid={!!error}
                  />
                  <Form.Control.Feedback type="invalid">{error}</Form.Control.Feedback>
                </Form.Group>
                <div className="d-grid">
                  <Button type="submit" variant="primary" size="lg" disabled={loading}>
                    {loading ? "Analyzing..." : "Analyze"}
                  </Button>
                </div>
              </Form>
            </Card.Body>
          </Card>
        </Col>
      </Row>

      {loading && (
        <div className="d-flex justify-content-center mt-4">
          <Loader />
        </div>
      )}

      {result && (
        <Row className="mt-5">
          <Col md={8} className="mx-auto">
            <Card className="shadow-sm">
              <Card.Header as="h4" className="bg-primary text-white">
                Analysis Results
              </Card.Header>
              <Card.Body>
                <Row>
                  <Col xs={12} md={6}>
                    <p>
                      <strong>Title:</strong> {result.title || "N/A"}
                    </p>
                    <p>
                      <strong>HTML Version:</strong> {result.html_version || "N/A"}
                    </p>
                  </Col>
                  <Col xs={12} md={6}>
                    <p>
                      <strong>Internal Links:</strong> {result.internal_links}
                    </p>
                    <p>
                      <strong>External Links:</strong> {result.external_links}
                    </p>
                    <p>
                      <strong>Inaccessible Links:</strong> {result.inaccessible_links}
                    </p>
                  </Col>
                </Row>
                <hr />
                <p>
                  <strong>Headings:</strong>{" "}
                  {result.headings ? JSON.stringify(result.headings) : "N/A"}
                </p>
                <p>
                  <strong>Login Form Present:</strong>{" "}
                  {result.has_login_form ? "Yes" : "No"}
                </p>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      )}
    </Container>
  );
};

export default HomePage;
