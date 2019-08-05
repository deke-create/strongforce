using System;
using System.Collections.Generic;
using StrongForce.Core.Exceptions;
using StrongForce.Core.Permissions;

namespace StrongForce.Core.Kits
{
	public class KitContract : Contract
	{
		public KitContract(Address address, Kit kit)
			: base(address)
		{
			this.Kit = kit;
			this.Acl.AddPermission(
				null,
				typeof(InstantiateKitAction),
				this.Address);
		}

		public bool Instantiated { get; private set; } = false;

		public Kit Kit { get; }

		protected override bool HandleAction(ActionContext context, Action action)
		{
			switch (action)
			{
				case InstantiateKitAction instantiateKitAction:
					this.HandleInstantiateKitAction(context, instantiateKitAction);
					return true;

				default:
					return base.HandleAction(context, action);
			}
		}

		private void HandleInstantiateKitAction(ActionContext context, InstantiateKitAction instantiateKitAction)
		{
			if (this.Instantiated)
			{
				throw new InvalidOperationException("Kit was already instantiated");
			}

			this.Instantiated = true;

			this.Kit.CreateContractHandler = this.CreateContract;
			this.Kit.SendActionHandler = this.SendAction;

			this.Kit.Instantiate(context.Sender);
		}
	}
}