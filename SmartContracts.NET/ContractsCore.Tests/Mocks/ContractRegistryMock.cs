﻿using System;
using System.Collections.Generic;
using System.Text;
using Action = ContractsCore.Actions.Action;

namespace ContractsCore.Tests.Mocks
{
	class ContractRegistryMock : ContractRegistry
	{
		public ContractRegistryMock(IAddressFactory addressFactory, object initialState = null)
			: base(addressFactory, initialState)
		{
		}

		public bool HandleSendAction(Action action, Address sender)
		{
			action.Origin = sender;
			action.Sender = sender;
			return this.HandleAction(action, action.Target);
		}
	}
}