import React from "react";
import Header from "./components/Header";
import Footer from "./components/Footer";
import HomePage from "./pages/HomePage";

const App = () => (
  <>
    <Header />
    <main className="min-vh-100 py-4">
      <HomePage />
    </main>
    <Footer />
  </>
);

export default App;
