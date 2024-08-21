import api from "../../api/api";
import { useEffect, useState } from "react";
import Datatable from "../../components/table/Datatable";
function Bill() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

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
      <Datatable data={expenses} type="expenses" />
    </div>
  );
}

export default Bill;
