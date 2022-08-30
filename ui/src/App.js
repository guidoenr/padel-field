import Header from './components/Header'
import Footer from './components/Footer'
import Reservation from './components/pages/Reservation';
import MyTurns from './components/pages/MyTurns';
import Login from './components/pages/Login'
import { Routes, Route } from "react-router-dom";

function App() {

  return (
    <div className="App bg-gradient-to-t from-[#8e9eab] to-base-100">
      <Header />
      <Routes>
        <Route path="/" element={<Reservation />} />
        <Route path="misturnos" element={<MyTurns />} />
        <Route path="login" element={<Login />} />
      </Routes>
      <Footer />
    </div>
  );
}

export default App;
