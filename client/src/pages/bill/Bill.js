import api from "../../api/api";
import { useEffect, useState } from "react";
import ExpensesTable from "../../components/table/ExpensesTable";
import ExpensesForm from "../../components/forms/ExpensesForm";
import { FaPlus } from "react-icons/fa";


function Bill() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [showForm, setShowForm] = useState(false);

  const handleOpenForm = () => setShowForm(true);
  const handleCloseForm = () => setShowForm(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await api.get('/expenses');
        setData(response.data);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;

  const expenses = data.expenses.length > 0 ? data.expenses : "Sem contas cadastradas"
  return (
    <div>
      <ExpensesTable data={expenses}/>
      <button className="add-button" onClick={handleOpenForm}>
        <FaPlus /> Adicionar Despesa
      </button>
      <div className="modal-form">
        {showForm && <ExpensesForm onClose={handleCloseForm} />}
      </div>
    </div>
  );
}

export default Bill;
