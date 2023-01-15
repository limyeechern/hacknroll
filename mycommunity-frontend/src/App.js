import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./Home/Home";
import Talk from "./Talk/Talk";
import Feed from "./Feed/Feed";
import About from "./About/About";
import { QueryClientProvider } from "react-query";
import { QueryClient } from "react-query";
import Helpline from "./Helpline/Helpline";


function App() {
  const queryClient = new QueryClient();
  return (
    <QueryClientProvider client={queryClient}>
      <div className="App">
        <div className="background"></div>
        <Routes>
          <Route exact path="/" element={<Home />} />
          <Route exact path="/talk" element={<Talk />} />
          <Route exact path="/feed" element={<Feed />} />
          <Route exact path="/about" element={<About />} />
          <Route exact path="/helpline" element={<Helpline />} />
        </Routes>
      </div>
    </QueryClientProvider>
  );
}

export default App;
