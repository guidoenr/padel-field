import Header from './components/Header'
import Reservation from './components/Reservation';
import Footer from './components/Footer'
import Login from './components/Login'
import { Routes, Route } from "react-router-dom";

function App() {

  return (
    <div className="App bg-gradient-to-t from-[#8e9eab] to-base-100">
      <Header />
      <Routes>
        <Route path="/" element={<Reservation />} />
        <Route path="login" element={<Login />} />
      </Routes>
      <Footer />
    </div>
  );
}

export default App;
