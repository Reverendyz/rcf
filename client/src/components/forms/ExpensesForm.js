import { useState } from 'react';
import Notification from '../notification/Notification';
import api from '../../api/api';

function ExpensesForm({ onClose }) {
  const [value, setValue] = useState('');
  const [description, setDescription] = useState('');
  const [type, setType] = useState('');
  const [status, setStatus] = useState(false);
  const [notification, setNotification] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      await api.post('/expenses/add', {
        value: parseFloat(value),
        description,
        type,
        status,
        is_active: true
      });
      setNotification('Expense added successfully');
      setValue('');
      setDescription('');
      setType('');
      setStatus(false);
    } catch (error) {
      setNotification('Failed to add expense');
      console.error('There was an error adding the expense!', error);
    }
  };

  return (
    <div className="floating-form-overlay">
      <div className="floating-form-container">
        <button type='button' className="floating-form-close" onClick={onClose}>Fechar Formulário</button>
        <h2>Adicionar Despesa</h2>
        <form onSubmit={handleSubmit}>
          <label>
            Valor:
            <input
              type="number"
              value={value}
              onChange={(e) => setValue(e.target.value)}
              required
            />
          </label>
          <label>
            Descrição:
            <input
              type="text"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              required
            />
          </label>
          <label>
            Tipo:
            <input
              type="text"
              value={type}
              onChange={(e) => setType(e.target.value)}
              required
            />
          </label>
          <label>
            Marcar como pago
            <input
              type="checkbox"
              checked={status}
              onChange={(e) => setStatus(e.target.checked)}
            />
          </label>
          <button type="submit">Adicionar</button>
        </form>
        {notification && (
          <Notification message={notification} onClose={() => setNotification('')} />
        )}
      </div>
    </div>
  );
}

export default ExpensesForm;
