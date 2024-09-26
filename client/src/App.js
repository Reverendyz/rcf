import './App.css';
import './index.css'
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Bill from './pages/bill/Bill';
import Account from './pages/account/Account';

function App() {
  return (
    <div className='page-container'>
      <Router>
        <div className="content">
          <nav>
            <Link to="/">
              <button className='nav-button'>Home</button>
            </Link>
            <Link to="/bill">
              <button className='nav-button'>Gerenciar Contas</button>
            </Link>
            <Link to="/account">
              <button className='nav-button'>Gerenciar Usuários</button>
            </Link>
          </nav>
          <div className='page'>
            <Routes>
              <Route path="/" />
              <Route path="/bill" element={<Bill />} />
              <Route path="/account" element={<Account />} />
            </Routes>
          </div>
        </div>
      </Router>
    </div>
  );
}

export default App;
