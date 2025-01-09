import React, { useState } from "react";
import { Container, Form, Button, Alert, Card } from "react-bootstrap";
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
      <h1 className="text-center mb-4">Web Page Analyzer</h1>
      <Form onSubmit={handleAnalyze}>
        <Form.Group>
          <Form.Label>Enter URL</Form.Label>
          <Form.Control
            type="text"
            placeholder="https://example.com"
            value={url}
            onChange={(e) => setURL(e.target.value)}
          />
        </Form.Group>
        <Button type="submit" variant="primary" className="mt-3">
          Analyze
        </Button>
      </Form>

      {loading && <Loader />}

      {error && <Alert variant="danger" className="mt-3">{error}</Alert>}

      {result && (
        <Card className="mt-4">
          <Card.Body>
            <h4>Analysis Results</h4>
            <p><strong>Title:</strong> {result.title}</p>
            <p><strong>HTML Version:</strong> {result.html_version}</p>
            <p><strong>Headings:</strong> {JSON.stringify(result.headings)}</p>
            <p><strong>Internal Links:</strong> {result.internal_links}</p>
            <p><strong>External Links:</strong> {result.external_links}</p>
            <p><strong>Inaccessible Links:</strong> {result.inaccessible_links}</p>
            <p><strong>Login Form Present:</strong> {result.has_login_form ? "Yes" : "No"}</p>
          </Card.Body>
        </Card>
      )}
    </Container>
  );
};

export default HomePage;
