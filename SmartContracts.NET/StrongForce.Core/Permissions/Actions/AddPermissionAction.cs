using System;
using System.Collections.Generic;

namespace StrongForce.Core.Permissions
{
	public class AddPermissionAction : Action
	{
		public AddPermissionAction(
			Address target,
			Permission permission,
			HashSet<Address> permittedAddress,
			Address receiver = null)
			: base(target)
		{
			this.PermittedAddress = permittedAddress;
			this.Receiver = receiver ?? target;
			this.Permission = permission;
		}

		public AddPermissionAction(
			Address target,
			Permission permission,
			Address permittedAddress,
			Address receiver = null)
			: base(target)
		{
			this.PermittedAddress = new HashSet<Address> { permittedAddress };
			this.Receiver = receiver ?? target;
			this.Permission = permission;
		}

		public HashSet<Address> PermittedAddress { get; }

		public Address Receiver { get; }

		public Permission Permission { get; }

		public static implicit operator AddPermissionAction(RemovePermittedAddressAction v)
		{
			throw new NotImplementedException();
		}
	}
}