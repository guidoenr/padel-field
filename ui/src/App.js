import Header from './components/Header'
import Footer from './components/Footer'
import Reservation from './components/pages/Reservation';
import MyTurns from './components/pages/MyTurns';
import Login from './components/pages/Login'
import Register from "./components/pages/Register";
import Profile from './components/pages/Profile'
import ConfirmReservation from './components/pages/ConfirmReservation';
import { Routes, Route, useLocation } from "react-router-dom";
import { AnimatePresence } from 'framer-motion';
import { useState } from "react";


function App() {
  const [emptyTurns, setEmptyTurns] = useState(true);
  const [nextTurns, setNextTurns] = useState(false);

  const location = useLocation()

  return (
    <div className="App bg-gradient-to-t from-[#8e9eab] to-base-100">
      <Header />
      <AnimatePresence>
        <Routes location={location} key={location.key}>
          <Route path="/" element={<Reservation />} />
          <Route path="confirmreservation" element={<ConfirmReservation />} />
          <Route path="misturnos" element={<MyTurns />} />
          <Route path="profile" element={<Profile />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
          <Route path="*" element={<Reservation />} />
        </Routes>
      </AnimatePresence>

      <Footer />
    </div>
  );
}

export default App;
