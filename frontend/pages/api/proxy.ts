import type {NextApiRequest, NextApiResponse} from 'next';
import {API_URL} from '../../config/api';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
	const {address} = req.query;

	try {
		const response = await fetch(
			`${API_URL}/wallet/transaction/history?address=${address}`
		);
		const data = await response.json();
		console.log('data in proxy', data);
		res.status(200).json(data);
	} catch (error) {
		console.log('error in proxy', error);
		res.status(500).json({error: 'Unable to fetch data'});
	}
}
