import http from 'k6/http';
import { check } from 'k6';

export const options = {
    scenarios: {
        create_orders: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s',
            duration: '30s',
            preAllocatedVUs: 200,
            maxVUs: 500,
        },
    },
};

const ORDER_SERVICE_URL = 'http://localhost:5959'

export default function() {
    let res = http.post(
        `${ORDER_SERVICE_URL}/orders`,
        JSON.stringify({
            productId: "719a2be1-67cb-4ffb-8ecc-ca718bb30ae0" // Change with correct product id
        }),
        {
            headers: { 'Content-Type': 'application/json' }
        });


    check(res, {
        'status is 200': (r) => r.status === 200,
    });
}
