import React, { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import Login from "./components/Login";
import Register from "./components/Register";
import Nav from "./components/Nav";
import Home from "./pages/Home";

export type User = {
  id: number;
  name: string;
  email: string;
};

function App() {
  const [user, setUser] = useState<User>();

  useEffect(() => {
    (async () => await getUser())();
  }, []);

  const getUser = async () => {
    const response = await fetch("/api/user/", {
      credentials: "include",
    });

    if (response.ok) {
      const content = await response.json();
      setUser(content);
    }
  }

  const handleLoggedIn = async () => { 
    await getUser() };

  return (
    <div className="App">
      <BrowserRouter>
        <Nav />

        <Routes>
          <Route path={"/"} element={<Home user={user} />} />
          <Route
            path={"/login"}
            element={<Login onLoggedIn={handleLoggedIn} />}
          />
          <Route path={"/register"} element={<Register />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
