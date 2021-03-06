import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import { LandingPage } from '../pages/LandingPage';
import { Home } from '../pages/Home';
import { NewSpectra } from '../pages/NewSpectra';
import { ListRequests } from '../pages/ListRequests';

export const AppRouter = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/home" element={<Home />} />
        <Route path="/new" element={<NewSpectra />} />
        <Route path="/list" element={<ListRequests />} />
      </Routes>
    </Router>
  )
}
